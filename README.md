# Autotwt
AutoTWT is a flexible social media automation tool that eliminates repetitive posting across multiple platforms. Write once, post everywhere.
## Pre-requisites
As of now this project only support `X(Twitter)`. We will add more platform in the future.

You need to have the X(Twitter) Credintials 
- Access Key
- Access Secret
- Consumer Key
- Consumer Secret

these 4 things you need it.
```sh
X_CONSUMER_KEY = xyz
X_CONSUMER_SECRET = xyz
X_ACCESS_TOKEN = xyz
X_ACCESS_SECRET = xyz
```


### what you will going to see upcoming furutre features
- Schedule you posts.
- Add Linkedin Platform 
## Build locally
If you want to Use it. you can use it locally on you system. by running the following commands.

Run the following command:

```sh
git clone "repository_url"
cd Autowtwt
```
Create a .env file in the Root directory 
```sh
touch .env
```
Run the Frontend.( it will run on localhost:3001)
```sh
pnpm run dev
```
Start the Backend
```sh
cd apps/backend/
air 
```

Now you can able to use it. 



<!-- ## What's inside?

This Turborepo includes the following packages/apps:

### Apps and Packages

- `docs`: a [Next.js](https://nextjs.org/) app with [Tailwind CSS](https://tailwindcss.com/)
- `web`: another [Next.js](https://nextjs.org/) app with [Tailwind CSS](https://tailwindcss.com/)
- `ui`: a stub React component library with [Tailwind CSS](https://tailwindcss.com/) shared by both `web` and `docs` applications
- `@repo/eslint-config`: `eslint` configurations (includes `eslint-config-next` and `eslint-config-prettier`)
- `@repo/typescript-config`: `tsconfig.json`s used throughout the monorepo

Each package/app is 100% [TypeScript](https://www.typescriptlang.org/).

### Building packages/ui

This example is set up to produce compiled styles for `ui` components into the `dist` directory. The component `.tsx` files are consumed by the Next.js apps directly using `transpilePackages` in `next.config.ts`. This was chosen for several reasons:

- Make sharing one `tailwind.config.ts` to apps and packages as easy as possible.
- Make package compilation simple by only depending on the Next.js Compiler and `tailwindcss`.
- Ensure Tailwind classes do not overwrite each other. The `ui` package uses a `ui-` prefix for it's classes.
- Maintain clear package export boundaries.

Another option is to consume `packages/ui` directly from source without building. If using this option, you will need to update the `tailwind.config.ts` in your apps to be aware of your package locations, so it can find all usages of the `tailwindcss` class names for CSS compilation.

For example, in [tailwind.config.ts](packages/tailwind-config/tailwind.config.ts):

```js
  content: [
    // app content
    `src/**/*.{js,ts,jsx,tsx}`,
    // include packages if not transpiling
    "../../packages/ui/*.{js,ts,jsx,tsx}",
  ],
```

If you choose this strategy, you can remove the `tailwindcss` and `autoprefixer` dependencies from the `ui` package.

### Utilities

This Turborepo has some additional tools already setup for you:

- [Tailwind CSS](https://tailwindcss.com/) for styles
- [TypeScript](https://www.typescriptlang.org/) for static type checking
- [ESLint](https://eslint.org/) for code linting
- [Prettier](https://prettier.io) for code formatting -->
