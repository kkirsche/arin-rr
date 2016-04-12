# Unofficial ARIN Route Registry Tool (arin-rr)
Tool to update the ARIN Route Registry. Templates used are based on:

https://www.arin.net/resources/routing/templates.html

## Installation

### From Source
```shell
go get -u github.com/kkirsche/arin-rr
```

### Configuration

Configuring arin-rr can be made easier via it's config file. This allows you to
set common values that will be re-used each time you use the tool unless overridden
on the CLI.

Example configuration:

```yaml
# $HOME/.arin-rr.yml
email:
  from: from@company.com
  smtp: smtp.company.com
arin:
  description: Example Company Network Services
  asn: 1234
  notify-email: notifyNetworking@company.com
  maintained-by: maintainer@company.com
  changed-email: changedByMe@company.com
```

### Example

#### Without a config file:

```shell
~/g/s/g/k/arin-rr git:master ❯❯❯ arin-rr route -r "1.2.3.4/24" -a 1234 -d "Example Company Network Services" -f "from@company.com" -n "notifyNetworking@company.com" -g "changedByMe@company.com" -m "maintainer@company.com"
To: rr@arin.net
From: from@company.com
Subject: route

route: 1.2.3.4/24
descr: This is a description
origin: AS1234
notify: notifyNetworking@company.com
mnt-by: maintainer@company.com
changed: changedByMe@company.com 20160406
source: ARIN
```

#### With the above config file:

```shell
~/g/s/g/k/arin-rr git:master ❯❯❯ arin-rr route -r "1.2.3.4/24"
To: rr@arin.net
From: from@company.com
Subject: route

route: 1.2.3.4/24
descr: This is a description
origin: AS1234
notify: notifyNetworking@company.com
mnt-by: maintainer@company.com
changed: changedByMe@company.com 20160406
source: ARIN
```
