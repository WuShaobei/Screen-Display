'''
Author: WuShaobei
Date: 2022-10-13 15:58:49
LastEditTime: 2022-10-13 16:09:14
FilePath: /Screen-Display/test.py
Description: 
'''
from locust import HttpUser, TaskSet, task

# 定义用户行为
class UserBehavior(TaskSet):

    @task
    def baidu_index(self):
        self.client.get("/")

class WebsiteUser(HttpUser):
    task = [UserBehavior] # 指向一个定义的用户行为类
    min_wait = 3000 # 执行事务之间用户等待时间的下界（单位：毫秒）
    max_wait = 6000 # 执行事务之间用户等待时间的上界（单位：毫秒）

# locust -f  test.py --host="http://127.0.0.1:1432/api/data/getAllDataByFromFundingStatistic"