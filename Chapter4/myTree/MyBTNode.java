package myTree;
import java.util.*;

public class MyBTNode {
	public int data;
	public MyBTNode left;
	public MyBTNode right;
	public MyBTNode parent;

	public MyBTNode() {
		data = 0;
		left = null;
		right = null;
		parent = null;
	}

	public MyBTNode(int d) {
		data = d;
		left = null;
		right = null;
		parent = null;
	}

	public MyBTNode(int d, MyBTNode l, MyBTNode r) {
		data = d;
		left = l;
		right = r;
		parent = null;
		if(left != null) {
			left.parent = this;
		}
		if(right != null) {
			right.parent = this;
		}
	}

	public void setLeft(MyBTNode n) {
		if(n == null)
			return;
		left = n;	
		if(left != null) {
			left.parent = this;
		}	
	}

	public void setRight(MyBTNode n) {
		if(n == null)
			return;
		right = n;	
		if(right != null) {
			right.parent = this;
		}
	}

	public int getHeight() {
		int numL = left != null ? left.getHeight() : 0;
		int numR = right != null ? right.getHeight() : 0;
		return Math.max(numL, numR);
	}
	
	public void insert(MyBTNode n) {
		if( n == null) 
			return;
		if(left == null)
			left = n;	
		else if(right == null) 
			right = n;
		else			
			left.insert(n);		

	}

	public Boolean isBST() {
		if(left != null) {
			if((left.data > data) || (!left.isBST())) {
				return false;
			}
		}
		if(right != null) {
			if((right.data <= data) || (!right.isBST())) {
				return false;
			}
		}		
		return true;		
		
	}

	public void buildBSTInOrder(MyBTNode n) {
		if(n == null)
			return;
		if(n.data <= data) 
			if(left == null)
				left = n;
			else
				left.buildBSTInOrder(n);
		else if(n.data > data)
			if(right == null)
				right = n;
			else
				right.buildBSTInOrder(n);		
	}
}









