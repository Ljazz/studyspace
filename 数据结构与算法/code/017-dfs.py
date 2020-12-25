# 最大油田代码
def MaxAreaOflsland(grid):    # grid为二维数组，其中存储地理信息
    row = len(grid)     # row记录二维数组的行数，也是地图的y轴长度
    col = len(grid[0])  # col记录二维数组的列数，也是地图的x轴长度
    arrived = [[False for j in range(col)] for i in range(row)]     # arrived为一个二维数组，存储一块土地是否被访问
    ans = 0     # 记录油田最大面积

    # 深度优先代码
    def DFS(x, y):    # grid为二维数组，其中存储地理信息
        if 0 <= x < row and 0 <= y < col and not arrived[x][y] and grid[x][y] == 1:
            arrived[x][y] = True
            return 1 + DFS(x-1, y) + DFS(x+1, y)+DFS(x, y-1) + DFS(x, y+1)
        else:
            return 0

    for i in range(row):
        for j in range(col):
            area = DFS(i, j)
            if area > ans:
                ans = area
    return ans


if __name__ == "__main__":
    grid = [
        [0, 0, 0, 0, 1, 1, 0],
        [0, 1, 1, 0, 1, 1, 0],
        [0, 1, 1, 0, 0, 0, 0],
        [0, 0, 1, 0, 0, 1, 1],
        [0, 0, 0, 0, 0, 0, 0],
        [0, 0, 1, 1, 0, 0, 0],
        [0, 0, 0, 1, 0, 0, 1],
    ]
    ret = MaxAreaOflsland(grid)
    print(ret)
