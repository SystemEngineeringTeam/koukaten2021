import datetime
import json
import os
import re
import subprocess
from time import sleep

import requests
import schedule

# カメラ計測間隔
IGNITION_INTERVAL = 1
# APIサーバホスト名
API_HOST = "http://localhost:8081"

# タイムゾーン
JST = datetime.timezone(datetime.timedelta(hours=9))

CAMERA_SERVER_STR="camera-server | "
def camera_server_print(text):
    print(CAMERA_SERVER_STR + text)


def json_serial(obj):
    if isinstance(obj, datetime.datetime):
        return obj.isoformat()
    

def camera_task():
    now = datetime.datetime.now(JST)
    print("カメラ起動")
    command = ['python3', os.path.join(os.path.dirname(__file__), 'detect.py')]
    res = subprocess.run(command, stdout=subprocess.PIPE)
    if res.returncode != 0:
        message = "ERROR: Failed to run detect.py."
        camera_server_print(message)
        return
    # post処理
    res = int(re.sub(r"\D", "", res.stdout.decode()))
    response = requests.post(f'{API_HOST}/people', data=json.dumps({'people': res, 'datetime': f'{now:%Y-%m-%d %H:%M:00}'}, default=json_serial))
    print(response.status_code)
    

for minute in range(0, 60, IGNITION_INTERVAL):
    schedule.every().hours.at(f':{minute:02}').do(camera_task)

camera_server_print("camera server serving...")
while True:
    schedule.run_pending()
    sleep(1)
