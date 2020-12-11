class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None


class BST:
    def __init__(self, tlist):
        self.root = TreeNode(tlist[0])
        for i in tlist[1:]:
            self.insert(i)

    def preorder(self, node):   # 先序遍历
        if node is None:
            return
        print(node.val)
        self.preorder(node.left)
        self.preorder(node.right)

    def inorder(self, node):   # 中序遍历
        if node is None:
            return
        self.preorder(node.left)
        print(node.val)
        self.preorder(node.right)

    def postorder(self, node):   # 后序遍历
        if node is None:
            return
        self.preorder(node.left)
        self.preorder(node.right)
        print(node.val)

    def insert(tlist):
        cnt = 1
        tree = [0]
        for item in tlist:
            tree.append(TreeNode(item))
        for item in tlist:
            if item.val == 'null':  # 若节点为 'null' 则不加入tree中
                continue
            if 2 * cnt < len(tlist) and tree[2*cnt].val != 'null':  # 找到每个节点的左子节点
                item.left = tree[2*cnt]
            if (2 * cnt + 1) < len(tlist) and tree[2*cnt+1].val != 'null':   # 找到每个节点的右子节点
                item.right = tree[2*cnt+1]
            cnt += 1
