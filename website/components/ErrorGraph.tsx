import { Button, Center, Stack, Title } from "@mantine/core";
import React from "react";
import { Refresh } from "tabler-icons-react";
import useGraph from "../hooks/useGraph";

type ErrorGraphProps = {
  errorMessage: string;
};

export default function ErrorGraph({ errorMessage }: ErrorGraphProps) {
  const { refetch } = useGraph();

  return (
    <Center style={{ flex: 4, height: "100%" }} inline>
      <Stack align={"center"}>
        <Title order={1}>{errorMessage}</Title>
        <Button
          radius={"xl"}
          size="xl"
          color={"teal"}
          variant="subtle"
          onClick={() => refetch()}
        >
          <Refresh size={"48px"} />
        </Button>
      </Stack>
    </Center>
  );
}
