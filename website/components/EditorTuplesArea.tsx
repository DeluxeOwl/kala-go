import { Stack } from "@mantine/core";

type EditorTuplesAreaProps = {
  children?: React.ReactNode;
};

const EditorTuplesArea = ({ children }: EditorTuplesAreaProps) => {
  return (
    <Stack
      style={{ flex: 4, height: "100%", borderRight: "1px solid gray" }}
      spacing={0}
    >
      {children}
    </Stack>
  );
};

export default EditorTuplesArea;
