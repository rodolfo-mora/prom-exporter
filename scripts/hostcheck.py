#!/usr/bin/env python

import json
import requests

def hostcheck(beat):
    requests.post("http://localhost:8080/api/v1/hostcheck", data=json.dumps(beat))

if __name__ == "__main__":
    for i in range(100):
        beat = {
            "name": "host%d.something.com"
        }
        beat['name'] = beat['name'] % i
        hostcheck(beat)
    print("Done posting")
