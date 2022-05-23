import { Kbd, List, Modal } from "@mantine/core";
import { useDisclosure, useHotkeys } from "@mantine/hooks";
import type { NextPage } from "next";
import { useEffect } from "react";
import EditorArea from "../components/EditorArea";
import EditorTuplesArea from "../components/EditorTuplesArea";
import GraphArea from "../components/GraphArea";
import Ide from "../components/Ide";
import TuplesArea from "../components/TuplesArea";
import { showHelpNotif } from "../util/notifications";

const Home: NextPage = () => {
  const [editorOpen, handlers] = useDisclosure(true);
  const [modalOpen, modalHandlers] = useDisclosure(false);

  useHotkeys([
    ["mod+M", () => handlers.toggle()],
    [
      "mod+shift+k",
      () => {
        modalHandlers.toggle();
      },
    ],
  ]);

  useEffect(() => {
    const timeoutId = setTimeout(() => {
      showHelpNotif();
    }, 1000);

    return () => clearTimeout(timeoutId);
  }, []);

  return (
    <>
      <Modal opened={modalOpen} onClose={() => modalHandlers.close()}>
        <List>
          <List.Item>
            <Kbd>Ctrl</Kbd> + <Kbd>Shift</Kbd> + <Kbd>K</Kbd> - bring up this
            menu
          </List.Item>
          <List.Item>
            <Kbd>Ctrl</Kbd> + <Kbd>K</Kbd> - reload changed from the editor
          </List.Item>
          <List.Item>
            <Kbd>Ctrl</Kbd> + <Kbd>M</Kbd> - make graph fullscreen
          </List.Item>
        </List>
      </Modal>
      <Ide>
        {editorOpen && (
          <EditorTuplesArea>
            <EditorArea>Editor</EditorArea>
            <TuplesArea>Tuples</TuplesArea>
          </EditorTuplesArea>
        )}
        <GraphArea>Graph</GraphArea>
      </Ide>
    </>
  );
};

export default Home;
