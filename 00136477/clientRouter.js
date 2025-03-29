let routes = null;
let routeContainer = null;

function processRoutes() {
  const currentRoute = window.location.pathname;

  if (!routeContainer) return;

  if (routes[currentRoute]) {
    routeContainer.innerHTML = routes[currentRoute]();
  } else {
    for (const route in routes) {
      const paramNames = [];
      const routeRegex = new RegExp(
        "^" +
          route.replace(/:([^/]+)/g, (_, paramName) => {
            paramNames.push(paramName);
            return "([^/]+)";
          }) +
          "$"
      );

      const match = currentRoute.match(routeRegex);
      if (match) {
        const params = paramNames.reduce((acc, paramName, index) => {
          acc[paramName] = match[index + 1];
          return acc;
        }, {});

        routeContainer.innerHTML = routes[route](params);
        handleLinks();
        return;
      }
    }

    routeContainer.innerHTML = routes[404]();
  }

  handleLinks();
}
function handleLinks() {
  const links = document.querySelectorAll("a[data-href]");
  links.forEach((link) => {
    link.addEventListener("click", handleLinkClick);
  });
}

function handleLinkClick(e) {
  e.preventDefault();
  const href = e.currentTarget.dataset.href;

  if (href) {
    history.pushState(null, "", href);
    processRoutes();
  }
}

function handleRouteChange() {
  window.onpopstate = processRoutes;
}

export function initializeRouter(routeList, container) {
  routes = routeList;
  routeContainer = container;

  processRoutes();
  handleRouteChange();
}
