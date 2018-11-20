# Definition for a binary tree node.
class TreeNode
    attr_accessor :val, :left, :right
    def initialize(val)
        @val = val
        @left, @right = nil, nil
    end
end

def sum_numbers(root)
    sum_helper(root, "", 0)
end

def sum_helper(n, v, sum)
    return 0 unless n
    unless n.left || n.right
       return sum += ((v+n.val.to_s).to_i)
    end

    lsum = sum_helper(n.left, v + n.val.to_s, sum)
    rsum = sum_helper(n.right, v + n.val.to_s, sum)
    
    lsum + rsum
end