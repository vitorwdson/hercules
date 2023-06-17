import * as trpcNext from "@trpc/server/adapters/next";
import { appRouter } from "../../../../server/routers/_app";

// export API handler
// @see https://trpc.io/docs/api-handler
const handler = trpcNext.createNextApiHandler({
  router: appRouter,
  createContext: () => ({}),
});

export { handler as GET, handler as POST };
