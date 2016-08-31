import myGraph.*;
import java.util.*;
public class FindRoute {


	public Boolean findRoute(Graph g, GraphNode n, GraphNode d) {
		Map visit = g.bfs(n);
		if(visit.containsKey(d) && (Boolean)visit.get(d) == true) {
			return true;
		}
		else {
			return false;
		}
	}

	public static void main(String[] args) {
		Graph g = new Graph();
		FindRoute f = new FindRoute();
		GraphNode origin = new GraphNode("A");
		g.setOrigin(origin);
		//Test 1
		origin.setNeighbor(new GraphNode("B"));
		origin.setNeighbor(new GraphNode("C"));
		g.getNodeFromVertex("C").setNeighbor(new GraphNode("D"));
		g.getNodeFromVertex("D").setNeighbor(new GraphNode("E"));
		Boolean result = f.findRoute(g, origin, g.getNodeFromVertex("G"));
		
		//Test 2
		Boolean result = f.findRoute(g, origin, g.getNodeFromVertex("E"));
		System.out.println(result);

	}
}


