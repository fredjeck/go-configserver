import { PropsWithChildren } from "react";

export default function Main(props:PropsWithChildren){
    return (
        <main className="isolate bg-white px-6 py-24 sm:py-32 lg:px-8">
            <div
        className="absolute inset-x-0 -z-10 transform-gpu overflow-hidden blur-3xl sm:top-[-10rem]"
        aria-hidden="true"
      >
        <div
          className="relative left-1/2 -z-10 aspect-[1155/678] w-[25rem] max-w-none bg-gradient-to-r from-green-300 via-blue-500 to-purple-600 opacity-30"
        />
      </div>
          {props.children}
      </main>
    );
}