import { Box } from "@mantine/core";

type TuplesAreaProps = {
  children?: React.ReactNode;
};

const TuplesArea = ({ children }: TuplesAreaProps) => {
  return (
    <Box
      style={{
        flex: 1,
      }}
    >
      {children}
    </Box>
  );
};

export default TuplesArea;
