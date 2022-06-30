const GetCookie = (name: string): string | undefined  => {
  const matches: RegExpMatchArray = document.cookie.match(new RegExp(
    `(?:^|; )${name.replace(/([.$?*|{}()[]\\\/+^])/g, '\\$1')}=([^;]*)`,
  ));
  return matches ? decodeURIComponent(matches[1]) : undefined;
};

export default GetCookie;
