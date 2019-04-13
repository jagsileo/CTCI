import myTree.*;
import java.util.*;

public class MyBT {
	private int MAX_TREE_SIZE = 10;
	private BTNode root;
	private BTNode[] children;
	private int nodeCount;


	public MyBT() {
		root = null;
		nodeCount = 0;
	}

	public MyBT(BTNode n) {
		root = n;
		nodeCount = 0;
	}

	public int getNodeCount() {
		return children.length;
	}

	public BTNode[] getChildren() {
		return children;
	}

	public BTNode getRoot() {
		return root;
	}

	public Boolean setRoot(BTNode n) {
		if(n == null) {
			return false;
		}

		if(root != null) {
			root = n;
		}
		return (root != null);
	}

	public void addChild(BTNode n) {
		if(nodeCount = 0) {
			root = n;
		}
		
		else if(nodeCount < MAX_TREE_SIZE) {
			BTNode last = children[nodeCount - 1];
			last.setLeft(n);
			children[nodeCount] = n;
			nodeCount++;
		}
		else {
			System.out.println("Tree is filled");
		}
	}


}






