import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AddOrganazerComponent } from './add-organazer.component';

describe('AddOrganazerComponent', () => {
  let component: AddOrganazerComponent;
  let fixture: ComponentFixture<AddOrganazerComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [AddOrganazerComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AddOrganazerComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
