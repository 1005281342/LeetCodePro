# -*- coding:utf-8 -*-
import os, sys, json, shutil
import urllib.request
from bs4 import BeautifulSoup

if __name__ == '__main__':
    print("#" * 10)
    os.chdir(os.path.dirname(os.path.realpath(__file__)))
    # index = int(sys.argv[1])
    index = sys.argv[1]
    url = r'https://leetcode-cn.com/api/problems/all/'
    headers = {
        'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:47.0) Gecko/20100101 Firefox/47.0',
        'Content-Type': 'application/json'
    }
    req = urllib.request.Request(url=url, headers=headers)
    res = json.loads(urllib.request.urlopen(req).read().decode('utf-8'))
    print(res['user_name'])
    for obj in res['stat_status_pairs']:
        if obj['stat']['frontend_question_id'] != index:
            continue
        title = obj['stat']['question__title']
        slug = obj['stat']['question__title_slug']
        # sub = os.path.join('src', 'p%04d' % index)
        sub = os.path.join('src', 'p%04d' % int(index))
        print("start creat")
        if not os.path.exists(sub):
            os.mkdir(sub)
            print("creat dir")
        attr = {
            'id': index,
            'title': title,
            'url': 'https://leetcode-cn.com/problems/%s/' % slug
        }
        fo = open(os.path.join(sub, 'attr.json'), 'w')
        json.dump(attr, fo, indent=4)
        print(sub)
