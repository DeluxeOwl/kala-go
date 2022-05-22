import { Box } from "@mantine/core";

type EditorAreaProps = {
  children?: React.ReactNode;
};

const EditorArea = ({ children }: EditorAreaProps) => {
  return (
    <Box
      style={{
        border: "1px solid red",
        flex: 1,
      }}
    >
      {children}
    </Box>
  );
};

export default EditorArea;
