/*
三角形也有费马点，三角形费马点是这样定义的：寻找三角形内的一个点，使得三个顶点到该点的距离之和最小。
三角形费马点的做法是：
（1）若有一个角大于120度，那么这个角所在的点就是费马点。
（2）若不存在，那么对于三角形ABC，任取两条边（假设AB、AC），向外做等边三角形得到C' 和 A'  ，那么AA' 和CC' 的交点就是费马点。

那么对于这题n多边形，我采取的策略完全不同，采用了模拟退火的做法，这种做法相对比较简单，也可以用在求三角形费马点之中。
模拟退火算法就跟数值算法里面的自适应算法相同的道理
（1）定义好步长
（2）寻找一个起点，往8个方向的按这个步长搜索。
（3）如果找到比答案更小的点，那么以这个新的点为起点，重复（2）
（4）如果找不到比答案更小的点，那么步长减半，再尝试（2）
（5）直到步长小于要求的答案的精度就停止。
*/


#include <iostream>
#include <cmath>
#include <cstdio>
using namespace std;

const int offx[]={1,1,0,-1,-1,-1,0,1};
const int offy[]={0,1,1,1,0,-1,-1,-1};
const int maxn=110;
const double eps=1e-6;

struct point{
    double x,y;
    point(double x0=0,double y0=0){x=x0;y=y0;}
    double getdis(point p){
        return sqrt((x-p.x)*(x-p.x)+(y-p.y)*(y-p.y));
    }
}p[maxn];

int n;

double getsum(point p0){
    double sum=0;
    for(int i=0;i<n;i++) sum+=p0.getdis(p[i]);
    return sum;
}

class Solution {
public:
    double getMinDistSum(vector<vector<int>>& positions) {
        n = positions.size();
        for(int i=0;i<n;i++) {
            p[i].x = positions[i][0];
            p[i].y = positions[i][1];
        }
        double ans=getsum(p[0]),step=100;
        point s=p[0],d;
        while(step>eps){
            bool flag=false;
            for(int i=0;i<8;i++){
                d=point(s.x+step*offx[i],s.y+step*offy[i]);
                double tmp=getsum(d);
                if(tmp<ans){
                    s=d;
                    ans=tmp;
                    flag=true;
                }
            }
            if(!flag) step/=2;
        }
        return ans;
    }
};