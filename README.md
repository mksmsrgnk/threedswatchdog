# watchdog
- Build app
- Create config file:
```
SERVICE_URL="<URL To Check>"
WORKER_SERVICE="<service name>"
KANNEL_ADDRESS="<kanel sms-box:port>"
KANNEL_USERNAME="<username>"
KANNEL_PASSWORD="<password>"
SMS_FROM="<outgoing number>"
SMS_TO="<destination number1,destination number2>"
```
- Add cron job:
```
* * * * * source /home/user/watchdog/config; /home/user/watchdog/watchdog >> /home/user/watchdog/watchdog.log 2>&1
```
 