"""
给定一个初始的路径
获取路径下的文件夹列表dirs并遍历得到dir_name
    如果dir_name是.git，则移除
    如果不是则放入队列中
"""
import os
import shutil
from collections import deque

start = "/Users/v_jesonou/Desktop/baselib_help/src/github.com"


def do():
    dq = deque()
    dq.append(start)
    while dq:
        name = dq.popleft()
        dq.extend(get_no_git_dir_list(name))


def get_no_git_dir_list(name):
    dir_list = list()
    for path in os.listdir(name):
        if not os.path.isdir(name + "/" + path):
            continue
        if path == ".git":
            # 删除文件夹
            shutil.rmtree(name + "/" + path)
        else:
            dir_list.append(name + "/" + path)
    return dir_list


if __name__ == '__main__':
    do()
