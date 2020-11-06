import { CloudMusicPage } from './app.po';

describe('cloud-music App', function() {
  let page: CloudMusicPage;

  beforeEach(() => {
    page = new CloudMusicPage();
  });

  it('should display message saying app works', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('app works!');
  });
});
