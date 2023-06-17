"use client";

import { api } from "../utils/api";

export default function Page() {
  const { data } = api.hello.useQuery({
    text: "Test",
  });

  return <h1>{data?.greeting}</h1>;
}
