#!/usr/bin/env python3

from genericpath import exists
import json
import os
import glob

SOURCE_CONF_TEMP_DIR = "build/temp"
TARGET_CONF_DIR = "build"

final = "{}"
output_list=[]

logging = {"critical": True, "debug": True, "error": True, "warn": True, "mqtt": False}

final_obj = json.loads(final)
final_obj["logging"] = logging


for filename in glob.glob(SOURCE_CONF_TEMP_DIR +"/*.json"):
    file = open(filename)
    obj_dict = json.load(file)
    file.close
    for key in obj_dict.keys():
        value = obj_dict[key]
        # print(key+" -> "+str(value))
        keylist = final_obj.keys()
        if key in keylist:
            final_obj[key] += value
        else:
            final_obj[key] = value

# print(json.dumps(final_obj, indent=2))
targetfile = open(TARGET_CONF_DIR+"/config.json.tmpl","w+")
targetfile.write(json.dumps(final_obj, indent=2))
targetfile.close()
