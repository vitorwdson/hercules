import NextAuth from "next-auth";
import { authOptions } from "../../../../server/auth";

// TODO: Add appropriate type to next auth handler.
// eslint-disable-next-line
const handler = NextAuth(authOptions);

export { handler as GET, handler as POST };
