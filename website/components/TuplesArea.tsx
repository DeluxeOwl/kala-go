import { Box } from "@mantine/core";

type TuplesAreaProps = {
  children?: React.ReactNode;
};

const TuplesArea = ({ children }: TuplesAreaProps) => {
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

export default TuplesArea;
