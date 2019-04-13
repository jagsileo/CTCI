package myTree;
import java.util.*;
public class BTNode{
	private int data;
	private BTNode left;
	private BTNode right;
	//	private BTNode parent;
	
	public BTNode() {
		data = 0;
		left = null;
		right = null;
	}

	public BTNode(int d) {
		data = d;
		left = null;
		right = null;
		//parent = null;
	}

	public BTNode(int d, BTNode l, BTNode r) {
		data = d;
		left = l;
		right = r;
		//		left.setParent(this);
		//		right.setParent(this);
	}

	public BTNode getLeft() {
		if(left != null)
			return left;
		else
			return null;
	}


	public BTNode getRight() {
		if(right != null)
			return right;
		else
			return null;
	}


	public int getData() {
		return data;
	}

	public Boolean setLeft(BTNode a) {
		if(a == null) {
			return false;
		}
		if(left == null) {
			left = a;
			//left.setParent(a);
		}
		return (left != null);
	}


	public Boolean setRight(BTNode a) {
		if(a == null) {
			return false;
		}
		if(right == null) {
			right = a;
			//right.setParent(a);
		}
		return (right != null);
	}


	public Boolean setData(int a) {
		data = a;
		return true;
	}

	public void printBTNode() {
		if(this != null) {
			System.out.println(this.getData());
		}
		else {
			System.out.println("Empty");
		}
		if(this.getLeft() != null) {
			this.getLeft().printBTNode();			
		}
		if(this.getRight() != null) {
			this.getRight().printBTNode();
		}
	}

	public int depthOfBT() {
		int dl=0, dr = 0;		
		if(this.getLeft() != null)
			dl = this.getLeft().depthOfBT();
		if(this.getRight() != null)
			dr = this.getRight().depthOfBT();
		int max = (dl >= dr) ? dl : dr;
		return max + 1;

	}


}



















