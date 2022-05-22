import type { NextPage } from "next";
import EditorArea from "../components/EditorArea";
import EditorTuplesArea from "../components/EditorTuplesArea";
import GraphArea from "../components/GraphArea";
import Ide from "../components/Ide";
import TuplesArea from "../components/TuplesArea";

const Home: NextPage = () => {
  return (
    <Ide>
      <EditorTuplesArea>
        <EditorArea>Editor</EditorArea>
        <TuplesArea>Tuples</TuplesArea>
      </EditorTuplesArea>
      <GraphArea>Graph</GraphArea>
    </Ide>
  );
};

export default Home;
