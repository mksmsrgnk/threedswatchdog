# watchdog
- Build app
- Create config file:
```
export SERVICE_URL="<URL To Check>"
export WORKER_SERVICE="<service name>"
export KANNEL_ADDRESS="<kanel sms-box:port>"
export KANNEL_USERNAME="<username>"
export KANNEL_PASSWORD="<password>"
export SMS_FROM="<outgoing number>"
export SMS_TO="<destination number1,destination number2>"
```
- Add cron job:
```
* * * * * source /home/user/watchdog/config; /home/user/watchdog/watchdog >> /home/user/watchdog/watchdog.log 2>&1
```
 