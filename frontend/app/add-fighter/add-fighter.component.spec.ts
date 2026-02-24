import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddFighterComponent } from './add-fighter.component';

describe('AddFighterComponent', () => {
  let component: AddFighterComponent;
  let fixture: ComponentFixture<AddFighterComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [AddFighterComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AddFighterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
