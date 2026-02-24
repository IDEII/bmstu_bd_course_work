import { ComponentFixture, TestBed } from '@angular/core/testing';

import { InvateOnTournamentComponent } from './invate-on-tournament.component';

describe('InvateOnTournamentComponent', () => {
  let component: InvateOnTournamentComponent;
  let fixture: ComponentFixture<InvateOnTournamentComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [InvateOnTournamentComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(InvateOnTournamentComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
