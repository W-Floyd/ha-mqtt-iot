# Documentation 

## Basic script

```
$ cat time.sh

#!/bin/bash

### DEPENDENCIES_BEGIN
# date
### DEPENDENCIES_END

### CONFIG_TEMPLATE_BEGIN
# {
#     "sensor": 
#     [
#         {
#             "name": "___HOSTNAME___ Time",
#             "object_id": "___HOSTNAME___-time-sensor",
#             "unique_id": "___HOSTNAME___-time-sensor",
#             "icon": "mdi:clock-time-five-outline",
#             "state": [
#                 "___SCRIPTS_DIR______SCRIPT_NAME___"                 
#             ],
#             "mqtt": {
#                 "update_interval": 10
#             }
#         }
#     ]
# }
### CONFIG_TEMPLATE_END

date +'%Y-%m-%d %H:%M:%S'

```
This script shows a simple time sensor. It shows the time of the system, where ha-mqtt-iot is installed on in Home Assistant.

### Basic Structure: Dependencies

Follow the script above. 

In the following predefined section the dependencies for this script are defiend. So the dependency of the script above is the "date" command (Please take care about the # and the whitespaces). 

```
### DEPENDENCIES_BEGIN
# date
### DEPENDENCIES_END
```

Add more dependencies seperated by white spaces. During the configuration generation the dependcies listed here are where checked by the command "which". If one dependency is not met, the configuration for this script is ignored. 

### Basic Structure: Configuration Template 

This is the heart of the script and ties ha-mqtt-iot to this script. Between the following predfiend section you define your Home Assistant entity which is populated by ha-mqtt-iot. 

```
### CONFIG_TEMPLATE_BEGIN
# {
#     "sensor": 
#     [
#         {
#             "name": "___HOSTNAME___ Time",
#             "object_id": "___HOSTNAME___-time-sensor",
#             "unique_id": "___HOSTNAME___-time-sensor",
#             "icon": "mdi:clock-time-five-outline",
#             "state": [
#                 "___SCRIPTS_DIR______SCRIPT_NAME___"                 
#             ],
#             "mqtt": {
#                 "update_interval": 10
#             }
#         }
#     ]
# }
### CONFIG_TEMPLATE_END
```

This config template does nothing else than calling the script. Here are some predfiend variables you have <b>NOT</b> to care about. 

* ```___HOSTNAME___```-> is substitudet during the configuration build with you hostname
* ```___SCRIPT_DIR___``` -> is substituted during the configuration build with the target directory e.g. if you install ha-mqtt-iot on your system (```make install```) this variable is substituted with /home/\<your-username>/.config/ha-mqtt-iot/scripts/ automatically. If you make a debug build (```make debug```) it is replaced by \<your-current-directory>/build/ automatically 
* ```___SCRIPT_NAME___``` -> is substituted with the name of the script dynamically and automatically. So if you name the script ```date.sh``` it will be replaced with with ```date.sh```. 

### Basic Structure: The Rest 

Following the example above this is the rest. 

```
date +'%Y-%m-%d %H:%M:%S'
```

This is your script. This command line is executed when mqtt sends the sensor value to mqtt. 

## Advanced Scripting 

### A Varible

Ok - you like to go crazy with scripting. We heard you. Lets introduce an other simple script:

```
#!/bin/bash

### DEPENDENCIES_BEGIN
# ip jq
### DEPENDENCIES_END

### CONFIG_TEMPLATE_BEGIN
# {
#     "sensor": 
#     [
#         {
#             "name": "___HOSTNAME___ ___INTERFACE___ IP Address",
#             "object_id": "___HOSTNAME___-___INTERFACE___-ip-address",
#             "unique_id": "___HOSTNAME___-___INTERFACE___-ip-address",
#             "icon": "mdi:ip-network",
#             "state": [
#                 "___SCRIPTS_DIR______SCRIPT_NAME___",
#                 "___INTERFACE___"
#             ],
#             "mqtt": {
#                 "update_interval": 10
#             }
#         }
#     ]
# }
### CONFIG_TEMPLATE_END

### HELPER_BEGIN
# { "___INTERFACE___": "eth0" }
### HELPER_END

__interface="${1}"
__state=$(cat "/sys/class/net/${__interface}/operstate")

if [ "${__state}" = 'down' ]; then
    echo "${__state}"
else
    ip a show "${__interface}" | grep inet | grep -v inet6 | awk '{print $2}' | awk -F "/" '{print $1}'
fi

```

In the above configuration template you see the freely defined variable ```___INTERFACE___```. This variable is defiend in the helper section below the configuration template section 

```
### HELPER_BEGIN
# { "___INTERFACE___": "eth0" }
### HELPER_END
```

During the generation of the final configuration for ha-mqtt-iot the variable ```___INTERFACE___``` is replaced by ```eth0```. 

So fine.

### Varible in a loop

Writing a script for every network interface in your system is <i>boring</i>....

Lets introduce a loop: 

```
### HELPER_BEGIN
# { "___INTERFACE___": [ "eth0", "wlan0", "dmz0" ] }
### HELPER_END
```

Now during the make process the automatic config generator creates 3 configurations. The first by replacing ```___INTERFACE___``` with ```eth0```, the scond with ```wlan0```, and the final one with ```dmz0```. 

### Variable replaced in a dynamic loop (scripting !!! :-))

Ok - sometimes you simply do not know how many interfaces (and what are their names) are available on the target system. Lets make it more dynamic. 

```
### HELPER_BEGIN
# { "___INTERFACE___“ : {"shell": "ls /sys/class/net/ | grep -v docker0 | grep -v lo" } }
### HELPER_END
```

The shell command ```ls /sys/class/net/``` returns a list of every available interface on the target host. So the above script returns <i>every</i> interface <i>except</i> "docker0" and "lo". 

## Thinks to improve 

- [ ] More Variables - at the moment every script is tied to one variable only 
- [ ] Set Hardware (/dev/...) as a dependency - i don't need a battery sensor when the system has no battery 
- 

