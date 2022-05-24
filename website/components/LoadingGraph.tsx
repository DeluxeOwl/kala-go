import { Center, Loader, Stack, Title } from "@mantine/core";

export default function LoadingGraph() {
  return (
    <Center style={{ flex: 4, height: "100%" }} inline>
      <Stack align={"center"}>
        <Title order={1}>Loading</Title>
        <Loader size="xl" variant="dots" />
      </Stack>
    </Center>
  );
}
