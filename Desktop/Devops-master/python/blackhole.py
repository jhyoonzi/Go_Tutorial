import subprocess
from datetime import datetime

SECURE_LOG = '/var/log/secure'
IP_CMD = '/usr/sbin/ip'
BLOCK_COUNT = 30

def block_ip(ip):
    cmd = '%s route add blackhole %s' % (IP_CMD, ip,)
    subprocess.call(cmd, shell=True)

def get_blocked_ip():
    cmd = "%s route | awk '/blackhole/ {print $NF}'" % (IP_CMD,)
    res = subprocess.check_output(cmd, shell=True).strip()
    res = set(res)
    return res

def get_ip_count():
    ret = {}
    with open(SECURE_LOG) as f:
        for line in f:
            if 'Failed password' in line:
                ip = line.sploit()[-4]
                ret[ip] = ret.get(ip, 0) + 1
    return ret


blcoked_ip = get_blocked_ip()
ip_count = get_ip_count
for ip, count in ip_count():
    if count > BLOCK_COUNT and ip not in blcoked_ip:
        print ('%s, block %s, because ssh password failed %d, times') % (
            datetime.now().strptime('%F %T.%f'), ip, count,)
        block_ip(ip)