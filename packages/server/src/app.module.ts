import { ApolloDriver, ApolloDriverConfig } from '@nestjs/apollo';
import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { GraphQLModule } from '@nestjs/graphql';
import { AppResolver } from './app.resolver';
import { PrismaService } from '@bug-tracker/db';

@Module({
  imports: [
    ConfigModule.forRoot(),
    GraphQLModule.forRoot<ApolloDriverConfig>({
      driver: ApolloDriver,
      autoSchemaFile: 'schema.gql',
      playground: process.env.DEV_MODE == 'true',
      cors: {
        origin: process.env.FRONTEND_URL,
        credentials: true,
      },
    }),
  ],
  providers: [PrismaService, AppResolver],
})
export class AppModule {}
