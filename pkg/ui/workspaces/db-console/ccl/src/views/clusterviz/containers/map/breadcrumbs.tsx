// Copyright 2018 The Cockroach Authors.
//
// Licensed as a CockroachDB Enterprise file under the Cockroach Community
// License (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
//     https://github.com/cockroachdb/cockroach/blob/master/licenses/CCL.txt

import { util } from "@cockroachlabs/cluster-ui";
import clone from "lodash/clone";
import React from "react";
import { Link } from "react-router-dom";

import { LocalityTier } from "src/redux/localities";
import { CLUSTERVIZ_ROOT } from "src/routes/visualization";
import { generateLocalityRoute, getLocalityLabel } from "src/util/localities";
import { trustIcon } from "src/util/trust";

import mapPinIcon from "!!raw-loader!assets/mapPin.svg";

import "./breadcrumbs.styl";

interface BreadcrumbsProps {
  tiers: LocalityTier[];
}

const { intersperse } = util;

export class Breadcrumbs extends React.Component<BreadcrumbsProps> {
  render() {
    const paths = breadcrumbPaths(this.props.tiers);

    return (
      <div className="breadcrumbs">
        <span
          className="breadcrumbs__icon"
          dangerouslySetInnerHTML={trustIcon(mapPinIcon)}
        />
        {intersperse(
          paths.map((path, idx) => (
            <span key={idx}>
              {idx === paths.length - 1 ? (
                getLocalityLabel(path)
              ) : (
                <Link
                  to={CLUSTERVIZ_ROOT + generateLocalityRoute(path)}
                  className="breadcrumbs__link"
                >
                  {getLocalityLabel(path)}
                </Link>
              )}
            </span>
          )),
          <span className="breadcrumbs__separator"> &gt; </span>,
        )}
      </div>
    );
  }
}

function breadcrumbPaths(path: LocalityTier[]): LocalityTier[][] {
  const pathSoFar: LocalityTier[] = [];
  const output: LocalityTier[][] = [[]];
  path.forEach(tier => {
    pathSoFar.push(tier);
    output.push(clone(pathSoFar));
  });
  return output;
}
