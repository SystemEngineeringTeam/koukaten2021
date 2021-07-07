import numpy as np
import argparse
import time
from pathlib import Path

import cv2
import torch
import torch.backends.cudnn as cudnn

from models.experimental import attempt_load
from utils.datasets import LoadStreams, LoadImages, letterbox
from utils.general import check_img_size, check_requirements, check_imshow, non_max_suppression, apply_classifier, \
    scale_coords, xyxy2xywh, strip_optimizer, set_logging, increment_path, save_one_box
from utils.plots import colors, plot_one_box
from utils.torch_utils import select_device, load_classifier, time_synchronized


def detect(opt):
    # 引数を代入
    # source = 検出を行う画像，動画
    # weights = 検出に使用する重み
    # view_img = 出力結果を画面に表示するかどうか
    # save_txt = 出力結果をテキストファイルとして出力するかどうか
    # imgsz = 推測サイズ(ピクセル)
    source, weights, view_img, save_txt, imgsz = opt.source, opt.weights, opt.view_img, opt.save_txt, opt.img_size
    save_img = not opt.nosave and not source.endswith(
        '.txt')  # save inference images
    # ソースがウェブカメラであるかどうか
    # 今回のプロジェクトでは，これは必ずTrueになる．（source.isnumeric()はtrueであるため）
    webcam = source.isnumeric() or source.endswith('.txt') or source.lower().startswith(
        ('rtsp://', 'rtmp://', 'http://', 'https://'))

    # 出力ファイルを保存するためのパスを求めている
    save_dir = increment_path(
        Path(opt.project) / opt.name, exist_ok=opt.exist_ok)  # increment run
    (save_dir / 'labels' if save_txt else save_dir).mkdir(parents=True,
                                                          exist_ok=True)  # make dir

    # 初期化
    set_logging()
    device = select_device(opt.device)
    # 今回はcpuでの検出しか行わない予定なので，必ずFalseになる．
    half = device.type != 'cpu'  # half precision only supported on CUDA

    # モデルの読み込み
    model = attempt_load(weights, map_location=device)  # load FP32 model
    stride = int(model.stride.max())  # model stride
    imgsz = check_img_size(imgsz, s=stride)  # check img_size
    # モデルの名前を取得する(人間はnames[0])
    names = model.module.names if hasattr(
        model, 'module') else model.names
    if half:
        model.half()  # to FP16

    # Second-stage classifier
    classify = False
    # ここ，必ずFalseにならない？
    if classify:
        modelc = load_classifier(name='resnet101', n=2)  # initialize
        modelc.load_state_dict(torch.load(
            'weights/resnet101.pt', map_location=device)['model']).to(device).eval()

    # Set Dataloader
    vid_path, vid_writer = None, None
    # ここは必ずTrueになる
    if webcam:
        # imshowが可能かどうか判定する
        view_img = check_imshow()
        cudnn.benchmark = True  # set True to speed up constant image size inference
        # LoadStreamsクラスを初期化して受け取る
        dataset = LoadStreams(source, img_size=imgsz, stride=stride)
    else:
        dataset = LoadImages(source, img_size=imgsz, stride=stride)

    # 恐らく，必ずFalseになる．
    if device.type != 'cpu':
        model(torch.zeros(1, 3, imgsz, imgsz).to(device).type_as(
            next(model.parameters())))  # run once
    # 開始時間
    t0 = time.time()
    # データセットの数だけ繰り返す
    # LoadStreamsの__next__を参照
    # pathはsourceをclean_strにかけたもの,imgは画像のデータ，
    # im0sはimgの配列？, vid_capはNone
    for path, img, im0s, vid_cap in dataset:
        img = torch.from_numpy(img).to(device)
        img = img.half() if half else img.float()  # uint8 to fp16/32
        img /= 255.0  # 0 - 255 to 0.0 - 1.0
        if img.ndimension() == 3:
            img = img.unsqueeze(0)

        # Inference
        t1 = time_synchronized()
        pred = model(img, augment=opt.augment)[0]

        # Apply NMS
        # NMSは，同じクラスとして重複して認識された状態を抑制するアルゴリズム．
        # 例えば，同じ顔が3回認識されているといった状態を防ぐ．
        pred = non_max_suppression(
            pred, opt.conf_thres, opt.iou_thres, classes=opt.classes, agnostic=opt.agnostic_nms)
        t2 = time_synchronized()

        # Apply Classifier
        if classify:
            pred = apply_classifier(pred, modelc, img, im0s)

        # Process detections
        # for文だが，実験してみたところ，1フレームにつき1度しか回っていなかった．
        for i, det in enumerate(pred):  # detections per image
            # 以下，出力，保存用の文字列設定
            if webcam:  # batch_size >= 1
                p, s, im0, frame = path[i], f'{i}: ', im0s[i].copy(
                ), dataset.count
            else:
                p, s, im0, frame = path, '', im0s.copy(), getattr(dataset, 'frame', 0)

            p = Path(p)  # to Path
            save_path = str(save_dir / p.name)  # img.jpg
            txt_path = str(save_dir / 'labels' / p.stem) + \
                ('' if dataset.mode == 'image' else f'_{frame}')  # img.txt
            s += '%gx%g ' % img.shape[2:]  # print string
            # normalization gain whwh
            gn = torch.tensor(im0.shape)[[1, 0, 1, 0]]
            imc = im0.copy() if opt.save_crop else im0  # for opt.save_crop
            people = 0
            # detは実質的な検出結果
            if len(det):
                # Rescale boxes from img_size to im0 size
                det[:, :4] = scale_coords(
                    img.shape[2:], det[:, :4], im0.shape).round()

                # Print results
                for c in det[:, -1].unique():
                    n = (det[:, -1] == c).sum()  # detections per class
                    # add to string
                    s += f"{n} {names[int(c)]}{'s' * (n > 1)}, "
                    if c == 0:
                        people = int(n)

                # Write results
                for *xyxy, conf, cls in reversed(det):
                    if save_txt:  # Write to file
                        xywh = (xyxy2xywh(torch.tensor(xyxy).view(1, 4)
                                          ) / gn).view(-1).tolist()  # normalized xywh
                        # label format
                        line = (
                            cls, *xywh, conf) if opt.save_conf else (cls, *xywh)
                        with open(txt_path + '.txt', 'a') as f:
                            f.write(('%g ' * len(line)).rstrip() % line + '\n')

                    if save_img or opt.save_crop or view_img:  # Add bbox to image
                        c = int(cls)  # integer class
                        label = None if opt.hide_labels else (
                            names[c] if opt.hide_conf else f'{names[c]} {conf:.2f}')
                        plot_one_box(xyxy, im0, label=label, color=colors(
                            c, True), line_thickness=opt.line_thickness)
                        if opt.save_crop:
                            save_one_box(
                                xyxy, imc, file=save_dir / 'crops' / names[c] / f'{p.stem}.jpg', BGR=True)
            print(people)

            # Print time (inference + NMS)
            # print(f'{s}Done. ({t2 - t1:.3f}s)')

            # Stream results
            if view_img:
                cv2.imshow(str(p), im0)
                cv2.waitKey(1)  # 1 millisecond

            # Save results (image with detections)
            if save_img:
                # if source == '0':
                # save_path += '.jpg'
                # cv2.imwrite(save_path, im0)
                # return
                if dataset.mode == 'image':
                    cv2.imwrite(save_path, im0)
                else:  # 'video' or 'stream'
                    if vid_path != save_path:  # new video
                        vid_path = save_path
                        if isinstance(vid_writer, cv2.VideoWriter):
                            vid_writer.release()  # release previous video writer
                        if vid_cap:  # video
                            fps = vid_cap.get(cv2.CAP_PROP_FPS)
                            w = int(vid_cap.get(cv2.CAP_PROP_FRAME_WIDTH))
                            h = int(vid_cap.get(cv2.CAP_PROP_FRAME_HEIGHT))
                        else:  # stream
                            fps, w, h = 30, im0.shape[1], im0.shape[0]
                            save_path += '.mp4'
                        vid_writer = cv2.VideoWriter(
                            save_path, cv2.VideoWriter_fourcc(*'mp4v'), fps, (w, h))
                    vid_writer.write(im0)

    if save_txt or save_img:
        s = f"\n{len(list(save_dir.glob('labels/*.txt')))} labels saved to {save_dir / 'labels'}" if save_txt else ''
        # print(f"Results saved to {save_dir}{s}")

    # print(f'Done. ({time.time() - t0:.3f}s)')


def detect2(opt):
    save_dir = increment_path(
        Path(opt.project) / opt.name, exist_ok=opt.exist_ok)  # increment run
    (save_dir / 'labels' if opt.save_txt else save_dir).mkdir(parents=True,
                                                              exist_ok=True)  # make dir
    view_img = check_imshow()

    device = select_device(opt.device)
    model = attempt_load(opt.weights, map_location=device)  # load FP32 model

    names = model.module.names if hasattr(
        model, 'module') else model.names

    save_img = not opt.nosave and not opt.source.endswith('.txt')
    cap = cv2.VideoCapture(eval(opt.source) if opt.source.isnumeric() else opt.source)
    stride = int(model.stride.max())
    vid_path, vid_writer = None, None

    people = np.empty(0, dtype=int)

    for times in range(5):
        _, img0 = cap.read()
        # cv2.imshow('name', img0)
        # cv2.waitKey(1)

        # Letterbox
        img = letterbox(img0, opt.img_size, stride=stride)[0]

        # Stack
        img = np.stack(img, 0)

        # Convert
        img = img[:, :, ::-1].transpose(2, 0, 1)  # BGR to RGB, to 3x416x416
        img = np.ascontiguousarray(img)

        img = torch.from_numpy(img).to(device)
        img = img.float()  # uint8 to fp16/32
        img /= 255.0  # 0 - 255 to 0.0 - 1.0
        if img.ndimension() == 3:
            img = img.unsqueeze(0)

        pred = model(img, augment=opt.augment)[0]

        # Apply NMS
        # NMSは，同じクラスとして重複して認識された状態を抑制するアルゴリズム．
        # 例えば，同じ顔が3回認識されているといった状態を防ぐ．
        pred = non_max_suppression(
            pred, opt.conf_thres, opt.iou_thres, classes=opt.classes, agnostic=opt.agnostic_nms)

        # Apply Classifier
        # if classify:
        # pred = apply_classifier(pred, modelc, img, im0s)

        # Process detections
        det = pred[0]
        # 以下，出力，保存用の文字列設定
        im0, frame = img0.copy(), times
        
        p = Path(opt.source)  # to Path
        save_path = str(save_dir / p.name)  # img.jpg
        txt_path = str(save_dir / 'labels' / p.stem) + f'_{frame}'  # img.txt
        # normalization gain whwh
        gn = torch.tensor(im0.shape)[[1, 0, 1, 0]]
        imc = im0.copy() if opt.save_crop else im0  # for opt.save_crop
        # detは実質的な検出結果
        if len(det):
            # Rescale boxes from img_size to im0 size
            det[:, :4] = scale_coords(
                img.shape[2:], det[:, :4], im0.shape).round()

            people = np.append(people, int((det[:, -1] == 0).sum()))

            # Write results
            for *xyxy, conf, cls in reversed(det):
                if opt.save_txt:  # Write to file
                    xywh = (xyxy2xywh(torch.tensor(xyxy).view(1, 4)
                                      ) / gn).view(-1).tolist()  # normalized xywh
                    # label format
                    line = (
                        cls, *xywh, conf) if opt.save_conf else (cls, *xywh)
                    with open(txt_path + '.txt', 'a') as f:
                        f.write(('%g ' * len(line)).rstrip() % line + '\n')

                if save_img or opt.save_crop or view_img:  # Add bbox to image
                    c = int(cls)  # integer class
                    label = None if opt.hide_labels else (
                        names[c] if opt.hide_conf else f'{names[c]} {conf:.2f}')
                    plot_one_box(xyxy, im0, label=label, color=colors(
                        c, True), line_thickness=opt.line_thickness)
                    if opt.save_crop:
                        save_one_box(
                            xyxy, imc, file=save_dir / 'crops' / names[c] / f'{p.stem}.jpg', BGR=True)

    print(np.argmax(np.bincount(people)))

    # Print time (inference + NMS)
    # print(f'{s}Done. ({t2 - t1:.3f}s)')

    # Stream results
    if view_img:
        cv2.imshow(str(p), im0)
        cv2.waitKey(1)  # 1 millisecond

    # Save results (image with detections)
    if save_img:
        if vid_path != save_path:  # new video
            vid_path = save_path
            if isinstance(vid_writer, cv2.VideoWriter):
                vid_writer.release()  # release previous video writer

            fps, w, h = 30, im0.shape[1], im0.shape[0]
            save_path += '.mp4'
            vid_writer = cv2.VideoWriter(
                save_path, cv2.VideoWriter_fourcc(*'mp4v'), fps, (w, h))
        vid_writer.write(im0)

if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('--weights', nargs='+', type=str,
                        default='yolov5x6.pt', help='model.pt path(s)')
    # file/folder, 0 for webcam
    parser.add_argument('--source', type=str,
                        default='0', help='source')
    parser.add_argument('--img-size', type=int, default=640,
                        help='inference size (pixels)')
    parser.add_argument('--conf-thres', type=float,
                        default=0.25, help='object confidence threshold')
    parser.add_argument('--iou-thres', type=float,
                        default=0.45, help='IOU threshold for NMS')
    parser.add_argument('--device', default='',
                        help='cuda device, i.e. 0 or 0,1,2,3 or cpu')
    parser.add_argument('--view-img', action='store_true',
                        help='display results')
    parser.add_argument('--save-txt', action='store_true',
                        help='save results to *.txt')
    parser.add_argument('--save-conf', action='store_true',
                        help='save confidences in --save-txt labels')
    parser.add_argument('--save-crop', action='store_true',
                        help='save cropped prediction boxes')
    parser.add_argument('--nosave', action='store_true',
                        help='do not save images/videos')
    parser.add_argument('--classes', nargs='+', type=int,
                        help='filter by class: --class 0, or --class 0 2 3')
    parser.add_argument('--agnostic-nms', action='store_true',
                        help='class-agnostic NMS')
    parser.add_argument('--augment', action='store_true',
                        help='augmented inference')
    parser.add_argument('--update', action='store_true',
                        help='update all models')
    parser.add_argument('--project', default='runs/detect',
                        help='save results to project/name')
    parser.add_argument('--name', default='exp',
                        help='save results to project/name')
    parser.add_argument('--exist-ok', action='store_true',
                        help='existing project/name ok, do not increment')
    parser.add_argument('--line-thickness', default=3,
                        type=int, help='bounding box thickness (pixels)')
    parser.add_argument('--hide-labels', default=False,
                        action='store_true', help='hide labels')
    parser.add_argument('--hide-conf', default=False,
                        action='store_true', help='hide confidences')
    opt = parser.parse_args()
    check_requirements(exclude=('tensorboard', 'pycocotools', 'thop'))

    with torch.no_grad():
        if opt.update:  # update all models (to fix SourceChangeWarning)
            for opt.weights in ['yolov5s.pt', 'yolov5m.pt', 'yolov5l.pt', 'yolov5x.pt']:
                # detect(opt=opt)
                detect2(opt=opt)
                strip_optimizer(opt.weights)
        else:
            # detect(opt=opt)
            detect2(opt=opt)
