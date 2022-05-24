import { Button, Center, Stack, Title } from "@mantine/core";
import React from "react";
import { Refresh } from "tabler-icons-react";

type ErrorGraphProps = {
  errorMessage: string;
};

export default function ErrorGraph({ errorMessage }: ErrorGraphProps) {
  return (
    <Center style={{ flex: 4, height: "100%" }} inline>
      <Stack align={"center"}>
        <Title order={1}>{errorMessage}</Title>
        <Button radius={"xl"} size="xl" color={"teal"} variant="subtle">
          <Refresh size={"48px"} />
        </Button>
      </Stack>
    </Center>
  );
}
