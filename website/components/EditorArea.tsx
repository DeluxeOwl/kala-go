import { Box, useMantineColorScheme } from "@mantine/core";
import { useHotkeys, useViewportSize } from "@mantine/hooks";
import Editor from "@monaco-editor/react";
import { useEffect, useRef, useState } from "react";
import YAML from "yaml";
import { showError } from "../util/notifications";

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
  const [requestPayload, setRequestPayload] = useState<string>("");
  const { colorScheme } = useMantineColorScheme();
  const monacoRef = useRef(null);
  // Font for larger screens
  const { width } = useViewportSize();

  // Triggered each time I press Ctrl+K
  useEffect(() => {
    console.log(requestPayload);
  }, [requestPayload]);

  useEffect(() => {
    if (monacoRef) {
      // @ts-ignore
      monacoRef?.current?.layout({});
    }
  }, [width]);

  // For getting the value
  useHotkeys([["mod+K", () => handleEditorValue()]]);

  function handleEditorValue() {
    if (!monacoRef) {
      return;
    }
    // @ts-ignore
    const value: string = monacoRef?.current?.getValue()?.split("---");

    let requestPayload = new Array();
    for (const val of value) {
      try {
        requestPayload.push(YAML.parse(val));
      } catch (YAMLParseError) {
        // @ts-ignore
        const errorMessage = YAMLParseError?.toString();
        showError(errorMessage);

        return;
      }
    }

    setRequestPayload(JSON.stringify(requestPayload, null, 2));
  }

  function handleEditorWillMount(monaco: any) {
    // define custom theme
    monaco.editor.defineTheme("vs-dark-custom", {
      base: "vs-dark",
      inherit: true,
      rules: [],
      colors: {
        "editor.background": "#1a1b1e",
      },
    });
  }

  function handleEditorDidMount(editor: any, monaco: any) {
    monacoRef.current = editor;
  }

  return (
    <Box
      style={{
        flex: 1,
      }}
    >
      <Editor
        defaultLanguage="yaml"
        defaultValue={defaultExample}
        options={{
          minimap: {
            enabled: false,
          },
          fontSize: width > 2000 ? 20 : 16,
          wordWrap: "on",
        }}
        theme={colorScheme === "dark" ? "vs-dark-custom" : "vs-light"}
        beforeMount={handleEditorWillMount}
        onMount={handleEditorDidMount}
      />
    </Box>
  );
};

export default EditorArea;
