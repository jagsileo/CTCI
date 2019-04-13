import java.util.*;
public class GraphNode {
	public static String name;
	public static GraphNode[] neighbors;
	public static int MAX_SIZE = 20;
	public static Boolean visited;
	public static int count = 0;

	public GraphNode() {
		name = "";
		neighbors = new GraphNode[MAX_SIZE];
		visited = false;
		count = 0;
	}

	public GraphNode(String s) {
		name = s;
		visited = false;
		neighbors = new GraphNode[MAX_SIZE];
		count = 0;
	}
	

	public static int getNeighborsCount() {
		return count;
	}

	public static Boolean setNeighbor(GraphNode d) {
		if(count == MAX_SIZE) {
			System.out.println("Graph full");
			return false;
		}
		
		neighbors[count] = d;
		count++;
		return true;

	}
	
	public static Boolean hasNeighbor() {
		if(neighbors.length > 0)
			return true;
		else
			return false;
			
	}



}


