# -*- coding: utf-8 -*-
#!/usr/bin/env python3

from genericpath import exists
import json
import os
import glob


# Dealing with path-es - i hat path-es
abs_path = os.path.dirname(os.path.realpath(__file__))

rel_SOURCE_CONF_TEMP_DIR = "../build/temp"
rel_TARGET_CONF_DIR = "../build"

SOURCE_CONF_TEMP_DIR = os.path.join(abs_path, rel_SOURCE_CONF_TEMP_DIR)
TARGET_CONF_DIR = os.path.join(abs_path, rel_TARGET_CONF_DIR)

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
targetfile = open(TARGET_CONF_DIR+"/config.json.tmpl","w")
targetfile.write(json.dumps(final_obj, indent=4))
targetfile.close()
