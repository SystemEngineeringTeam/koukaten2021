from time import sleep

import schedule
import os
import subprocess

# カメラ計測間隔
IGNITION_INTERVAL = 2

CAMERA_SERVER_STR="camera-server | "
def camera_server_print(text):
    print(CAMERA_SERVER_STR + text)

def camera_task():
    print("カメラ起動")
    command = ['python3', os.path.join(os.path.dirname(__file__), 'detect.py')]
    res = subprocess.run(command, stdout=subprocess.PIPE)
    if res.returncode != 0:
        message = "ERROR: Failed to run detect.py."
        camera_server_print(message)
        return

    # post処理
    

    return
    
for minute in range(0, 60, IGNITION_INTERVAL):
    schedule.every().hours.at(f':{minute:02}').do(camera_task)

camera_server_print("camera server serving...")
while True:
    schedule.run_pending()
    sleep(1)
