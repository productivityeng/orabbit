import React from "react";
import styles from "./styles.module.css";

type FeatureItem = {
  title: string;
  Svg: React.ComponentType<React.ComponentProps<"svg">>;
  description: JSX.Element;
};

const FeatureList: FeatureItem[] = [
  {
    title: "Easy to Use",
    Svg: require("@site/static/img/undraw_docusaurus_mountain.svg").default,
    description: (
      <>
        Focus on large scale management with ease modifications preveting
        configurations that will broke your production environment
      </>
    ),
  },
  {
    title: "When, Who and Why",
    Svg: require("@site/static/img/undraw_docusaurus_tree.svg").default,
    description: (
      <>
        Don't waste time trying to figureout who made a modification in your broker. See an complete history
          of a specific artifact change, ordered by owner,author and reason
      </>
    ),
  },
  {
    title: "Prevent unalthorized modifications",
    Svg: require("@site/static/img/undraw_docusaurus_react.svg").default,
    description: (
      <>
        Enable approve condition to enforce a revision of a senior engineer for selected operations and artifacts
      </>
    ),
  },{
        title: "No longer error prone",
        Svg: require("@site/static/img/undraw_docusaurus_react.svg").default,
        description: (
            <>
                Enable approve condition to enforce a revision of a senior engineer for selected operations and artifacts
            </>
        ),
    },
];

function Feature({ title, Svg, description }: FeatureItem) {
  return (
    <div className="flex items-center shadow-md hover:shadow-xl rounded-lg p-10 transition transform ease-out duration-300 space-x-5">
      <div className="text--center">
        <Svg className={styles.featureSvg} role="img" />
      </div>
      <div className="">
        <h3 className="font-bold text-2xl ">{title}</h3>
        <p>{description}</p>
      </div>
    </div>
  );
}

export default function HomepageFeatures(): JSX.Element {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="grid grid-cols-1 xl:grid-cols-2 gap-4">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}
