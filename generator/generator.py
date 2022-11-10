# -*- coding: utf-8 -*-
#!/usr/bin/env python3

from genericpath import exists
import subprocess
import json
import os
import re
import sys
import glob

START_CONFIG = "### CONFIG_TEMPLATE_BEGIN"
END_CONFIG = "### CONFIG_TEMPLATE_END"
START_DEP = "### DEPENDENCIES_BEGIN"
END_DEP = "### DEPENDENCIES_END"
START_HELPER = "### HELPER_BEGIN"
END_HELPER = "### HELPER_END"



def find_between( s, first, last ):
    try:
        start = s.index( first ) + len( first )
        end = s.index( last, start )
        return s[start:end]
    except ValueError:
        return ""

def process_helper(mydict, key):
    if isinstance(mydict[key],str):
        resultlist = [mydict[key]]
    elif isinstance(mydict[key],dict):
        for nextkey in mydict[key]:
            if nextkey == "shell":
                result = subprocess.run(mydict[key][nextkey], stdout=subprocess.PIPE, shell=True)
                resultlist = result.stdout.decode("utf-8").split()
            else:
                resultlist = []
    elif isinstance(mydict[key],list):
        resultlist = mydict[key]
    else:
        resultlist = []
    return resultlist

def check_dependencies( dependencies_list ):
    ret = 0
    for dep in dependencies_list:
        cmd = "which "+dep
        result = subprocess.run(cmd.split(), stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)
        if result.returncode != 0:
            print("[ ERROR ] "+scriptname+": "+dep+" .... not installed")
            ret = 1
        else:
            print("[ INFO ] "+scriptname+": "+dep+" .... installed")
    return ret

def validateJSON( json_string ):
    try:
        json.loads(json_string)
    except ValueError as err:
        return False
    return True

def writeconfig( config_template, orig_scriptname, dest_filename):
    c1 = config_template.replace("___SCRIPT_NAME___",orig_scriptname)
    if validateJSON(c1) == False:
        print("[ ERROR ] "+orig_scriptname+": Substitution produces invalid configuration")
        return False
    json_obj = json.loads(c1)
    conf = open(dest_filename,"w", encoding='utf-8')
    conf.write(json.dumps(json_obj, indent=4))
    conf.close()
    return True


# Dealing with path-es - i hat path-es
abs_path = os.path.dirname(os.path.realpath(__file__))
rel_scripts_dir = "../scripts"
rel_target_conf_temp_dir = "../build/temp"

SCRIPT_DEV_DIR = os.path.join(abs_path, rel_scripts_dir)
TARGET_CONF_TEMP_DIR = os.path.join(abs_path, rel_target_conf_temp_dir)

if not os.path.exists(TARGET_CONF_TEMP_DIR):
    os.makedirs(TARGET_CONF_TEMP_DIR)

if not os.path.exists(SCRIPT_DEV_DIR):
    print("[ ERROR ] "+rel_scripts_dir+" does not exists")
    exit(0)

# print(SCRIPT_DEV_DIR)
# print(TARGET_CONF_TEMP_DIR)

# exit(0)

scripts = glob.glob(SCRIPT_DEV_DIR+"/*.sh")
for script in scripts:
    scriptname = script.split("/")[-1]
    print("[ INFO ] "+scriptname+" parsing ...")

    f = open(script, encoding='utf-8')
    data = f.read()
    f.close()
    
    # grap the dependency list from inside the script 
    # if one dependency is not met - the configuration is not created

    dependencies = find_between(data, START_DEP, END_DEP).replace("# ","").split()
    if check_dependencies(dependencies) != 0:
        ("[ ERROR ] "+scriptname+": dependencies not met")
        continue
    
    # get the configuration template from the script 
    # if none is found - the configuration is not created 
    config_template = find_between(data, START_CONFIG, END_CONFIG).replace("# ","")
    if config_template == "":
        print("[ ERROR ] "+scriptname+": configuration template not found in parsing ")
        continue
    else:
        print("[ INFO ] "+scriptname+": parsing configuration template")
    
    if validateJSON(config_template) == False:
        print("[ ERROR ] "+scriptname+": Substitution produces invalid configuration")
        continue

    # tricky part 
    # the helper 
    # sometimes you have to create a sensor or something else multiple times - this is were the helper comes into play 
    #
    # the helper consists of a key-value pair 
    # example:
    #
    # a) { "___CMD___": "sudo reboot" } 
    #    this is the easiest case - the key ___CMD___ in the configuration is replaced by "sudo reboot" - without the quotes
    #
    # b) { "___ITEM___": [ "alpha", "beta", "gamma" ] }
    #    in this case the configuration is created 3 times - one were ___ITEM___ is replaced with "alpha", in the second 
    #    ___ITEM___ is replaced by "beta", etc. 
    #    WARNING: the items must be elements of a list ( inside of square brackets )
    #
    # c) { "___INTERFACE___": { "shell": "ls /sys/class/net/" } }
    #    this is the "going crazy" case - this time the items where generated out of shell command. the generator (this script) 
    #    executes the command "ls /sys/class/net/" (which btw returns every network interface) and returns a list, which is substituted with the 
    #    ____INTERFACE___ key. Lets asume you have 2 network interfaces - eth0 and wlan0 -> the generator executes the script, gets the interfaces (eth0, wlan0),
    #    converts it to a list [ "eth0", "wlan0" ] and creates a configuration for every interface 
    
    ### multiple helpers ???
    helper = find_between(data, START_HELPER, END_HELPER).replace("# ","").replace("\n","")
    if helper != "":
        if validateJSON(helper) == False:
            print("[ ERROR ] "+scriptname+": Helper is not valid JSON")
            continue
        helper_dict = json.loads(helper)
        for key in helper_dict.keys():
            print(key)
            valuelist = process_helper(helper_dict,key)
            print(valuelist)
            for v in valuelist:
                c1 = config_template.replace(key,v)
                if validateJSON(c1) == False:
                    print("[ ERROR ] "+scriptname+": Substitution in helper ("+key+" -> "+v+") produces invalid configuration")
                    continue
                filename = TARGET_CONF_TEMP_DIR+"/"+scriptname.replace(".","-")+"-"+v+"-config.json"
                if not writeconfig(c1, scriptname, filename):
                    continue

    else:
        filename = TARGET_CONF_TEMP_DIR+"/"+scriptname.replace(".","-")+"-config.json"
        if not writeconfig(config_template, scriptname, filename):
            continue
        
    print("[ OK ] "+scriptname+": config successfully")

