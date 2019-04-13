import myTree.*;
import java.util.*;

public class CheckBalancedTree {

	public static int getHeight(BTNode n) {
		if(n == null)
			return -1;
		return Math.max(getHeight(n.getLeft()), getHeight(n.getRight()));	

	}

	public static Boolean checkBalTree(BTNode n) {

		if(n == null) {
			return true;
		}
		else {
			BTNode left = n.getLeft();
			BTNode right = n.getRight();
			int nLeft = 0, nRight = 0;
			nLeft = getHeight(left);
			nRight = getHeight(right);
			int diff = Math.abs(nLeft - nRight);
			if(diff > 1) {
				return false;
			}
			else {
				return checkBalTree(left) && checkBalTree(right);
			}
		}

	}

	public static void main(String[] args) {
		BTNode root = new BTNode(5);
		root.setLeft(new BTNode(3));
		root.setRight(new BTNode(7));
		root.getLeft().setLeft(new BTNode(2));
		root.getLeft().setRight(new BTNode(4));
		root.getLeft().getLeft().setRight(new BTNode(10));
		//root.getRight().setRight(new BTNode(8));
		//root.getRight().setRight(new BTNode(10));
		Boolean res = checkBalTree(root);
		System.out.println(res);

	}
}












