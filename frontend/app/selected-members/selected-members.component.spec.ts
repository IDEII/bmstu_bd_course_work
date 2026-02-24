import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SelectedMembersComponent } from './selected-members.component';

describe('SelectedMembersComponent', () => {
  let component: SelectedMembersComponent;
  let fixture: ComponentFixture<SelectedMembersComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [SelectedMembersComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SelectedMembersComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
