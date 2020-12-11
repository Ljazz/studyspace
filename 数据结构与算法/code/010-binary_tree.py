# 类的定义
class TreeNode:     # 二叉树节点的定义
    def __init__(self, val):
        self.val = val      # 二叉树的值
        self.left = None    # 左孩子节点
        self.right = None   # 右孩子节点


if __name__ == "__main__":
    Input = [0]     # Input列表用于存储输入
    tree = [0]      # tree列表用于存储节点
    Input = Input + input().split()
    cnt = 1
    for item in Input:
        tree.append(TreeNode(item))
    for item in tree[1:]:
        if item.val == 'null':  # 若节点为 'null' 则不加入tree中
            continue
        if 2 * cnt < len(Input) and tree[2*cnt].val != 'null':  # 找到每个节点的左子节点
            item.left = tree[2*cnt]
        if (2 * cnt + 1) < len(Input) and tree[2*cnt+1].val != 'null':   # 找到每个节点的右子节点
            item.right = tree[2*cnt+1]
        cnt += 1
