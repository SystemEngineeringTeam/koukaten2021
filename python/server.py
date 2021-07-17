import re
from wsgiref.simple_server import make_server

import json
import subprocess
import os

CAMERA_SERVER_STR="camera-server | "
def camera_server_print(text):
    print(CAMERA_SERVER_STR + text)

def app(environ, start_response):
    command = ['python3', 'detect.py']
    res = subprocess.run(command, stdout=subprocess.PIPE)
    if res.returncode != 0:
        message = "ERROR: Failed to run detect.py."
        status = '500 INTERNAL SERVER ERROR'
        headers = [
            ('Content-type', 'application/json; charset=utf-8'),
            ('Access-Control-Allow-Origin', '*')
        ]
        start_response(status, headers)
        camera_server_print(message)
        return [json.dumps({'status': status, 'message': message}).encode("utf-8")]

    res = int(re.sub(r"\D", "", res.stdout.decode()))
    status = '200 OK'
    headers = [
        ('Content-type', 'application/json; charset=utf-8'),
        ('Access-Control-Allow-Origin', '*'),
    ]
    start_response(status, headers)
    return [json.dumps({'people': res}).encode("utf-8")]


with make_server('', 8082, app) as httpd:
    camera_server_print("Serving on port 8082...")
    httpd.serve_forever()
