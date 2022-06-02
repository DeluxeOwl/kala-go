import { Box, useMantineColorScheme } from "@mantine/core";
import { useHotkeys, useViewportSize } from "@mantine/hooks";
import Editor from "@monaco-editor/react";
import { useEffect, useRef, useState } from "react";
import YAML from "yaml";
import useGraph from "../hooks/useGraph";
import useTuples from "../hooks/useTuples";
import {
  defaultPc,
  defaultTuples,
  defaultYaml,
  gdrivePc,
  gdriveTuples,
  gdriveYaml,
  rbacPc,
  rbacTuples,
  rbacYaml,
} from "../ideconfigs/configs";
import { fetchAll } from "../util/fetchAll";
import { showError, showSuccess } from "../util/notifications";
import ExampleChange from "./ExampleChange";

type EditorAreaProps = {
  children?: React.ReactNode;
};

type SelectValues =
  | ""
  | "Github"
  | "Your Authorization model"
  | "Documents"
  | "Google Drive"
  | "RBAC"
  | "Custom Roles";

const EditorArea = ({ children }: EditorAreaProps) => {
  const { colorScheme } = useMantineColorScheme();
  const monacoRef = useRef(null);
  // Font for larger screens
  const { width } = useViewportSize();
  const { refetch } = useGraph();
  const tuples = useTuples((s) => s.tuples);
  const subjects = useTuples((s) => s.getUniqueSubjects)();
  const setState = useTuples((s) => s.setState);
  const [selectValue, setSelectValue] = useState<SelectValues>("");

  useEffect(() => {
    if (monacoRef) {
      // @ts-ignore
      monacoRef?.current?.layout({});
    }
  }, [width]);

  // For getting the value
  useHotkeys([["mod+K", () => handleEditorValue()]]);

  useEffect(() => {
    switch (selectValue) {
      case "":
        break;
      case "Github":
        break;
      case "Your Authorization model":
        // @ts-ignore
        monacoRef?.current?.setValue(
          `
type: your_type
---
type: your_other_type
relations:
  your_rel: your_type
permissions:
  your_perm: your_rel
          `
        );
        setState([], []);
        break;
      case "Documents":
        // @ts-ignore
        monacoRef?.current?.setValue(defaultYaml);
        setState(defaultTuples, defaultPc);
        break;
      case "Google Drive":
        // @ts-ignore
        monacoRef?.current?.setValue(gdriveYaml);
        setState(gdriveTuples, gdrivePc);
        break;
      case "RBAC":
        // @ts-ignore
        monacoRef?.current?.setValue(rbacYaml);
        setState(rbacTuples, rbacPc);
        break;
      case "Custom Roles":
        // @ts-ignore
        monacoRef?.current?.setValue("");
        setState([], []);
        break;
      default:
        // @ts-ignore
        monacoRef?.current?.setValue(defaultYaml);
        setState(defaultTuples, defaultPc);
        break;
    }
  }, [selectValue]);

  const handleEditorValue = async () => {
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
    try {
      await fetchAll([requestPayload, subjects, tuples]);
      await refetch();
      showSuccess("Reloaded config");
    } catch (error) {
      // @ts-ignore
      showError(`in editor: ${error.message}`);
    }
  };

  const handleEditorWillMount = (monaco: any) => {
    // define custom theme
    monaco.editor.defineTheme("vs-dark-custom", {
      base: "vs-dark",
      inherit: true,
      rules: [],
      colors: {
        "editor.background": "#1a1b1e",
      },
    });
  };

  const handleEditorDidMount = (editor: any, monaco: any) => {
    monacoRef.current = editor;
  };

  return (
    <>
      <ExampleChange value={selectValue} setValue={setSelectValue} />
      <Box
        style={{
          flex: 1,
          minHeight: "50%",
        }}
      >
        <Editor
          defaultLanguage="yaml"
          defaultValue={defaultYaml}
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
    </>
  );
};

export default EditorArea;
