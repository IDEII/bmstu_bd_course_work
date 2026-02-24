import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RequestOnTournamentComponent } from './request-on-tournament.component';

describe('RequestOnTournamentComponent', () => {
  let component: RequestOnTournamentComponent;
  let fixture: ComponentFixture<RequestOnTournamentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [RequestOnTournamentComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RequestOnTournamentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
