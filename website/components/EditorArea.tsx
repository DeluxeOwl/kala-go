import { Box, useMantineColorScheme } from "@mantine/core";
import Editor, { useMonaco } from "@monaco-editor/react";
import { useEffect } from "react";

type EditorAreaProps = {
  children?: React.ReactNode;
};

const defaultExample = `
type: user
---
type: group
relations:
  member: user

---
type: folder
relations:
  reader: user | group#member

---
type: document
relations:
  parent_folder: folder
  writer: user
  reader: user
permissions:
  read: reader | writer | parent_folder.reader
  read_and_write: reader & writer
  read_only: reader & !writer
`;

const EditorArea = ({ children }: EditorAreaProps) => {
  const { colorScheme } = useMantineColorScheme();
  const monaco = useMonaco();

  useEffect(() => {
    if (monaco) {
      console.log("here is the monaco instance:", monaco);
    }
  }, [monaco]);

  return (
    <Box
      style={{
        flex: 1,
      }}
    >
      <Editor
        defaultLanguage="yaml"
        defaultValue={defaultExample}
        theme={colorScheme === "dark" ? "vs-dark" : "vs-light"}
      />
    </Box>
  );
};

export default EditorArea;
