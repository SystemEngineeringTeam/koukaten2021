import schedule
from time import sleep

# カメラ計測間隔
IGNITION_INTERVAL = 2

CAMERA_SERVER_STR="camera-server | "
def camera_server_print(text):
    print(CAMERA_SERVER_STR + text)

def camera_task():
    print("タスク実行中")
    
for minute in range(0, 60, IGNITION_INTERVAL):
    print(minute)
    schedule.every().hours.at(f':{minute:02}').do(camera_task)

while True:
    schedule.run_pending()
    sleep(1)