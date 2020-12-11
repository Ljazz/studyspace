# 树节点的定义
class TreeNode:
    def __init__(self, value):
        self.value = value
        self.left = None
        self.right = None


# 二叉搜索树定义
class BST:
    def __init__(self, tlist):
        self.root = TreeNode(tlist[0])      # 第一个元素建立为根节点
        for i in tlist[1:]:     # 按顺序将剩下元素插入二叉搜索树中
            self.insert(i)

    def search(self, node, parent, data):
        """
        在二叉搜索树中查找是否有键值为val的节点。
        : node: 当前节点
        : parent: 父亲节点
        : data: 要查询的值
        ：return：tuple类型(当前节点是否存在，节点本身，父节点)
        """
        if node is None:    # 当前节点为空，没有找到查找的元素
            return False, node, parent
        elif data == node.value:    # 查找到元素
            return True, node, parent
        elif data < node.value:     # 查找元素小于当前节点数据，进入左子树查找
            self.search(node.left, node, data)
        else:       # 查找元素大于当前节点数据，进入右子树查找
            self.search(node.right, node, data)

    def insert(self, data):
        """
        插入数据
        : data: 要插入的数据
        """
        # 查询要插入的数据是否已经存在
        exist, n, p = self.search(self.root, self.root, data)
        if exist:   # 数据节点已经存在，不需要插入
            return
        else:   # 不存在
            new_node = TreeNode(data)   # 新建节点
            if data > p.value:
                p.right = new_node
            else:
                p.left = new_node

    def getlast(self, node, data, maxn):
        """
        查找前驱
        """
        if node is None:    # 当前位置为空，仍没有找到data相等元素
            return False, maxn  # 返回目前及格过路径上最大的符合前驱要求的数
        elif data == node.value:    # 找到了与data相等的元素
            if node.left is None:   # 没有左子树
                return True, maxn
            else:   # 值为value的节点有左子树
                tmp = node.left
                while tmp.right is not None:
                    tmp = tmp.right
                return True, tmp    # 返回左子树中最大的数
        elif data < node.value:     # data小于当前节点数据，进入左子树
            self.getlast(node.left, data, maxn)
        else:   # data大于当前节点的数据，进入右子树
            if maxn.data < node.value:  # 如果当前数值已经比maxn中的前驱都大，更新maxn
                maxn = node
            return self.getlast(node.right, data, maxn)

    def getnext(self, node, data, minn):
        if node is None:    # 如果当前的位置为空，但仍没有查找到与data相等的元素
            return 0, minn  # 返回目前经过路径上最小的符合后继要求的数
        elif data == node.data:     # 找到了与data相等的元素
            if node.right is None:  # 值为data的节点没有右子树
                return 1, minn  # 返回目前经过路径上最小的符合后继要求的数
            else:   # 值为data的节点有右子树
                tmp = node.right    # 进入右子树后找到右子树中最小的数
                while tmp.left is not None:
                    tmp = tmp.left
                return 1, tmp   # 返回右子树中最小的数（最左边的数）
        elif data < node.data:  # data小于当前节点的数据，进入左子树
            if minn.data > node.data:  # 如果当前数据比已经存储在minn中的后继要小，更新minn
                minn = node
                return self.getlast(node.left, data, minn)
            else:   # data大于当前节点的数据，进入右子树
                return self.getlast(node.right, data, minn)

    def delete(self, root, data):
        """
        删除
        """
        exist, n, p = self.search(root, root, data)
        if not exist:   # 不存在
            print('data is not exist')
        else:
            if n.left is None:  # 要删除的节点左子树为空
                if n == p.left:     # 如果n是左孩子，把p的左孩子赋值为n的右孩子
                    p.left = n.right
                else:   # 如果n是右孩子，把p的右孩子复制为n的右孩子
                    p.right = n.right
                del n
            elif n.right is None:   # 要删除的节点右子树为空
                if n == p.left:     # 如果n是左孩子，把p的左孩子赋值为n的右孩子
                    p.left = n.left
                else:   # 如果n是右孩子，把p的右孩子复制为n的右孩子
                    p.right = n.left
                del n
            else:   # 左右子树都不为空
                tmp = n.right   # 进入n的右子树
                if tmp.left is None:    # 如果n的右孩子没有左子树，则n的右孩子就是n的后继
                    n.value = tmp.value
                    n.right = tmp.right
                else:
                    next = tmp.left
                    while next.left is not None:    # 在右子树中查找n的后继
                        tmp = next
                        next = next.left
                    n.value = next.value
                    tmp.left = next.right
                    del next
