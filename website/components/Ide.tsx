import { Group } from "@mantine/core";
import React from "react";

type IdeProps = {
  children?: React.ReactNode;
};

const Ide = ({ children }: IdeProps) => {
  return (
    <Group style={{ height: "100vh" }} spacing={0}>
      {children}
    </Group>
  );
};

export default Ide;
