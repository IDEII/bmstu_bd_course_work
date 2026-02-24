import { TestBed } from '@angular/core/testing';

import { OrganazerService } from './organazer.service';

describe('OrganazerService', () => {
  let service: OrganazerService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(OrganazerService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
