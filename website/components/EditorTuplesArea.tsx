import { Stack } from "@mantine/core";

type EditorTuplesAreaProps = {
  children?: React.ReactNode;
};

const EditorTuplesArea = ({ children }: EditorTuplesAreaProps) => {
  return (
    <Stack style={{ flex: 2, height: "100%" }} spacing={0}>
      {children}
    </Stack>
  );
};

export default EditorTuplesArea;
