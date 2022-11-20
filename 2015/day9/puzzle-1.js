const file = require("../../utils/file");
const path = require("path");

const contents = file.getWords(path.join(__dirname, "input.txt"), "\n", "");

console.log(contents);

const graph = contents.reduce((graph, line) => {
  const [from, _, to, __, distance] = line;

  if (from in graph) {
    graph[from][to] = Number(distance);
  } else {
    graph[from] = { [to]: Number(distance) };
  }

  if (to in graph) {
    graph[to][from] = Number(distance);
  } else {
    graph[to] = { [from]: Number(distance) };
  }

  return graph;
}, {});

console.log(graph);

const bruteForce = (start, graph, totalDistance = 0, visitedNodes = []) => {
  const visited2 = [...visitedNodes, start];

  let pendingNodes = Object.keys(graph).filter((k) => !visited2.includes(k));
  // console.log(
  //   `start`,
  //   start,
  //   `totalDistance`,
  //   totalDistance,
  //   `visited2`,
  //   visited2,
  //   `pendingNodes`,
  //   pendingNodes
  // );

  if (pendingNodes.length === 0) {
    return totalDistance;
  }

  const nextnodes = pendingNodes
    .filter((n) => graph[start] && graph[start][n])
    .map((n) => ({ next: n, distance: graph[start][n] }));

  //console.log(`nextnodes`, nextnodes);
  return Math.min(
    ...nextnodes.map((n) =>
      bruteForce(n.next, graph, totalDistance + n.distance, visited2)
    )
  );
};

const minDistance = Math.min(
  ...Object.keys(graph).map((k) => bruteForce(k, graph))
);

console.log(minDistance);
