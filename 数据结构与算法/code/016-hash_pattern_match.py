# 模式配置
def pingpong(games, players):
    game = games.split('，')
    player = players.split('，')
    if len(player) != len(game):    # 如果两个字符串的长度不一样，则肯定不匹配
        return False
    mydict = {}     # 记录模式字符串和目标字符串的对应关系
    used = {}   # 记录目前已经使用过的字符串
    for i in range(len(game)):
        if game[i] in mydict:
            if mydict[game[i]] != player[i]:    # 不是第一次出现，则检查映射关系是否一致
                return False
        else:
            if player[i] in used:   # 检查选手是否出现过，出现过不成立
                return False
            mydict[game[i]] = player[i]     # 第一次出现，则加入哈希表
            used[player[i]] = True  # 在used中保存已经使用过的。
    return True     # 没有任何问题，返回成立


if __name__ == "__main__":
    games = "男单，女双，女双，男单"
    players = "李四，张/王组合，张/王组合，李四"
    ret = pingpong(games, players)
    print(ret)
